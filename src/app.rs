use std::io;

use crossterm::event::{self, Event, KeyCode};
use tui::{backend::Backend, Terminal};
use tui_tree_widget::TreeItem;
use helper::StatefulTree;

use crate::{
    helper,
    mongo::{get_collections_from_db, get_database_names, get_users_from_db, get_views_from_db},
    ui::ui,
};

pub enum InputMode {
    Normal,
    Insert,
}

pub enum Focus {
    DatabaseBlock,
    InputBlock,
}

pub struct App<'a> {
    pub input: String,
    pub input_mode: InputMode,
    pub messages: Vec<String>,
    pub focus: Option<Focus>,
    mongo_uri: String,
    pub tree: StatefulTree<'a>,
}

impl<'a> Default for App<'a> {
    fn default() -> App<'a> {
        App {
            input: String::new(),
            input_mode: InputMode::Normal,
            messages: Vec::new(),
            focus: None,
            mongo_uri: String::new(),
            tree: StatefulTree::new(),
        }
    }
}

impl<'a> App<'a> {
    // TODO: return a Result
    pub fn populate_hashmap(&mut self) {
        get_database_names(self.mongo_uri.to_owned())
            .unwrap()
            .iter()
            .for_each(|database| {

                let mut database_items = Vec::new();
                let mut items: Vec<TreeItem> = Vec::new();
                get_collections_from_db(database.to_string(), self.mongo_uri.to_owned())
                    .unwrap()
                    .iter()
                    .for_each(|collection| {
                        items.push(TreeItem::new_leaf(collection.to_string()))
                    });
                database_items.push(TreeItem::new("Collections", items.clone()));

                items.clear();
                get_views_from_db(database.to_string(), self.mongo_uri.to_owned())
                    .unwrap()
                    .iter()
                    .for_each(|view| {
                        items.push(TreeItem::new_leaf(view.to_string()))
                    });
                database_items.push(TreeItem::new("Views", items.clone()));

                items.clear();
                get_users_from_db(database.to_string(), self.mongo_uri.to_owned())
                    .unwrap()
                    .iter()
                    .for_each(|user| {
                        items.push(TreeItem::new_leaf(user.to_string()))
                    });
                database_items.push(TreeItem::new("Users", items.clone()));

                self.tree.items.push(TreeItem::new(database.to_string(), database_items));
            });
    }

    // reset terminal before panic
    fn chain_hook(&mut self) {
        let original_hook = std::panic::take_hook();

        std::panic::set_hook(Box::new(move |panic| {
            helper::reset_terminal().unwrap();
            original_hook(panic);
        }))
    }
}

pub fn run_app<B: Backend>(terminal: &mut Terminal<B>, mut app: App) -> io::Result<()> {
    app.chain_hook();

    app.mongo_uri = std::env::var("MONGODB_URI").expect("Set MONGODB_URI variable!");

    app.populate_hashmap();

    loop {
        terminal.draw(|f| ui(f, &mut app))?;

        if let Event::Key(key) = event::read()? {
            match app.input_mode {
                InputMode::Normal => match key.code {
                    KeyCode::Char('i') => {
                        app.input_mode = InputMode::Insert;
                        app.focus = Some(Focus::InputBlock);
                    }
                    KeyCode::Char('d') => {
                        app.focus = Some(Focus::DatabaseBlock);
                    }
                    KeyCode::Char('q') => {
                        return Ok(());
                    }
                    KeyCode::Esc => {
                        app.focus = None;
                    }
                    KeyCode::Char('j') => match app.focus {
                        Some(Focus::DatabaseBlock) => {
                            app.tree.down();
                        }
                        _ => {}
                    },
                    KeyCode::Char('k') => match app.focus {
                        Some(Focus::DatabaseBlock) => {
                            app.tree.up();
                        }
                        _ => {}
                    },
                    KeyCode::Char('g') => match app.focus {
                        Some(Focus::DatabaseBlock) => {
                            app.tree.first();
                        }
                        _ => {}
                    },
                    KeyCode::Char('G') => match app.focus {
                        Some(Focus::DatabaseBlock) => {
                            app.tree.last();
                        }
                        _ => {}
                    },
                    KeyCode::Enter => match app.focus {
                        Some(Focus::DatabaseBlock) => {
                            app.tree.toggle();
                        }
                        _ => {}
                    },
                    KeyCode::Left => app.tree.left(),
                    KeyCode::Right => app.tree.right(),
                    _ => {}
                },
                InputMode::Insert => match key.code {
                    KeyCode::Enter => {
                        app.messages.push(app.input.drain(..).collect());
                    }
                    KeyCode::Char(c) => {
                        app.input.push(c);
                    }
                    KeyCode::Backspace => {
                        app.input.pop();
                    }
                    KeyCode::Esc => {
                        app.input_mode = InputMode::Normal;
                        app.focus = None;
                    }
                    _ => {}
                },
            }
        }
    }
}
