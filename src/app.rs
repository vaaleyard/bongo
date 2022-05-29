use std::io;

use crossterm::event::{self, Event, KeyCode};
use tui::{backend::Backend, Terminal};

use crate::ui::ui;

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
    pub database_selected: Option<usize>,
    pub database_tree_size: Option<usize>,
}

impl Default for App {
    fn default() -> App {
        App {
            input: String::new(),
            input_mode: InputMode::Normal,
            messages: Vec::new(),
            focus: None,
            database_selected: Some(0),
            database_tree_size: Some(0),
        }
    }
}

pub fn run_app<B: Backend>(terminal: &mut Terminal<B>, mut app: App) -> io::Result<()> {
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
