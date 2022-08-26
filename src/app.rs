use std::{collections::BTreeMap, io};

use crossterm::event::{self, Event, KeyCode};
use tui::{backend::Backend, Terminal};

use crate::{
    helper,
    mongo::{get_collections_from_db, get_database_names, get_users_from_db, get_views_from_db},
    tree::Database,
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

pub struct App {
    pub input: String,
    pub input_mode: InputMode,
    pub messages: Vec<String>,
    pub focus: Option<Focus>,
    // database block control
    pub is_menu: Vec<bool>,
    pub database_selected: Option<usize>,
    pub database_tree_size: Option<usize>,
    pub database_tree: BTreeMap<String, Database>,
    mongo_uri: String,
}

impl Default for App {
    fn default() -> App {
        App {
            input: String::new(),
            input_mode: InputMode::Normal,
            messages: Vec::new(),
            is_menu: Vec::new(),
            focus: None,
            database_selected: Some(0),
            database_tree_size: Some(0),
            database_tree: BTreeMap::new(),
            mongo_uri: String::new(),
        }
    }
}

impl App {
    // TODO: return a Result
    pub fn populate_hashmap(&mut self) {
        get_database_names(self.mongo_uri.to_owned())
            .unwrap()
            .iter()
            .for_each(|database| {
                let mut database_object = Database::default();

                get_collections_from_db(database.to_string(), self.mongo_uri.to_owned())
                    .unwrap()
                    .iter()
                    .for_each(|collection| {
                        database_object.new_collection(collection.to_string());
                    });

                get_views_from_db(database.to_string(), self.mongo_uri.to_owned())
                    .unwrap()
                    .iter()
                    .for_each(|view| {
                        database_object.new_view(view.to_string());
                    });

                get_users_from_db(database.to_string(), self.mongo_uri.to_owned())
                    .unwrap()
                    .iter()
                    .for_each(|user| {
                        database_object.new_user(user.to_string());
                    });

                self.database_tree
                    .entry(database.to_string())
                    .or_insert(database_object);
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
                            // I confess it's weird, but it seems to be working
                            if app.database_tree_size
                                > usize::checked_add(
                                    app.database_selected.unwrap(),
                                    usize::try_from(1).unwrap(),
                                )
                            {
                                app.database_selected = usize::checked_add(
                                    app.database_selected.unwrap(),
                                    usize::try_from(1).unwrap(),
                                );
                            }
                        }
                        _ => {}
                    },
                    KeyCode::Char('k') => match app.focus {
                        Some(Focus::DatabaseBlock) => {
                            if app.database_selected != Some(0) {
                                app.database_selected = usize::checked_sub(
                                    app.database_selected.unwrap(),
                                    usize::try_from(1).unwrap(),
                                );
                            }
                        }
                        _ => {}
                    },
                    KeyCode::Char('g') => match app.focus {
                        Some(Focus::DatabaseBlock) => {
                            app.database_selected = Some(0);
                        }
                        _ => {}
                    },
                    KeyCode::Char('G') => match app.focus {
                        Some(Focus::DatabaseBlock) => {
                            app.database_selected =
                                usize::checked_sub(app.database_tree_size.unwrap(), 1);
                        }
                        _ => {}
                    },
                    KeyCode::Enter => match app.focus {
                        Some(Focus::DatabaseBlock) => {
                            app.messages
                                .push(app.is_menu[app.database_selected.unwrap()].to_string());
                        }
                        _ => {}
                    },
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
