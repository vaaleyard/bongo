use tui::{
    backend::Backend,
    layout::{Alignment, Constraint, Direction, Layout, Rect},
    style::{Color, Modifier, Style},
    text::{Span, Spans, Text},
    widgets::{Block, BorderType, Borders, List, ListItem, ListState, Paragraph},
    Frame,
};
use unicode_width::UnicodeWidthStr;

use crate::app::{App, Focus, InputMode};

pub fn ui<B: Backend>(f: &mut Frame<B>, app: &mut App) {
    let chunks = Layout::default()
        .margin(1)
        .constraints([Constraint::Length(3), Constraint::Min(0)].as_ref())
        .split(f.size());

    draw_input(f, app, chunks[0]);
    {
        let chunks = Layout::default()
            .constraints([Constraint::Percentage(30), Constraint::Percentage(80)].as_ref())
            .direction(Direction::Horizontal)
            .split(chunks[1]);

        draw_database_tree(f, app, chunks[0]);
        draw_preview(f, app, chunks[1]);
    }
}

fn draw_input<B: Backend>(f: &mut Frame<B>, app: &App, area: Rect) {
    let block = Block::default()
        .borders(Borders::ALL)
        .border_type(BorderType::Rounded)
        .title(Span::raw("Input"));
    // .title(Span::styled("Input", Style::default().fg(Color::Gray)));

    let (input_normal_mode_message, input_normal_mode_style) = (
        vec![
            Span::styled("Press ", Style::default().fg(Color::DarkGray)),
            Span::styled("q", Style::default().add_modifier(Modifier::BOLD)),
            Span::styled(" to exit, ", Style::default().fg(Color::DarkGray)),
            Span::styled("i", Style::default().add_modifier(Modifier::BOLD)),
            Span::styled(" to insert.", Style::default().fg(Color::DarkGray)),
        ],
        Style::default().add_modifier(Modifier::RAPID_BLINK),
    );

    let mut text = Text::from(Spans::from(input_normal_mode_message));
    text.patch_style(input_normal_mode_style);

    let input = match app.input_mode {
        InputMode::Insert => Paragraph::new(app.input.as_ref())
            .style(
                Style::default()
                    .add_modifier(Modifier::BOLD)
                    .fg(Color::Cyan),
            )
            .block(block),
        InputMode::Normal => Paragraph::new(text).style(Style::default()).block(block),
    };
    // TODO: better insert text formatting (add margin to input field,
    //       like a space before the word "Press")
    f.render_widget(input, area);

    match app.input_mode {
        InputMode::Normal => {}
        InputMode::Insert => {
            f.set_cursor(
                // cursor at the the end of the typing text
                area.x + app.input.width() as u16 + 1,
                area.y + 1,
            )
        }
    }
}

fn draw_database_tree<B: Backend>(f: &mut Frame<B>, app: &mut App, area: Rect) {
    let chunks = Layout::default()
        .constraints([Constraint::Percentage(100)])
        .direction(Direction::Vertical)
        .split(area);

    {
        let chunks = Layout::default()
            .constraints([Constraint::Percentage(100)])
            .direction(Direction::Vertical)
            .split(chunks[0]);

        let items = [
            ListItem::new("> admin"),
            ListItem::new("> config"),
            ListItem::new("> local"),
        ];
        app.database_tree_size = Some(items.len());

        let databases = List::new(items)
            .block(
                Block::default()
                    .title("Databases")
                    .title_alignment(Alignment::Center)
                    .borders(Borders::ALL)
                    .border_type(BorderType::Rounded),
            )
            .style(match app.focus {
                Some(Focus::DatabaseBlock) => Style::default().fg(Color::Cyan),
                _ => Style::default().fg(Color::White),
            })
            .highlight_style(Style::default().add_modifier(Modifier::ITALIC))
            .highlight_symbol(">");

        let mut state = ListState::default();
        state.select(app.database_selected);
        f.render_stateful_widget(databases, chunks[0], &mut state);
    }
}

fn draw_preview<B: Backend>(f: &mut Frame<B>, app: &App, area: Rect) {
    let chunks = Layout::default()
        .direction(Direction::Vertical)
        .constraints([Constraint::Percentage(100)])
        .split(area);

    let block = Block::default()
        .title("Preview")
        .title_alignment(Alignment::Center)
        .borders(Borders::ALL)
        .border_style(Style::default().fg(Color::White))
        .border_type(BorderType::Rounded)
        .style(Style::default().bg(Color::Black));

    let command: Vec<ListItem> = app
        .messages
        .iter()
        .enumerate()
        .map(|(_, message)| {
            let content = vec![Spans::from(Span::raw(format!("{}", message)))];
            ListItem::new(content)
        })
        .collect();

    let command = List::new(command).block(block);
    f.render_widget(command, chunks[0]);
}

// pub fn print_events() -> crossterm::Result<()> {
//     // `read()` blocks until an `Event` is available
//     match read()? {
//         Event::Key(event) => match event.code {
//             KeyCode::Tab => println!("This is a tab"),
//             _ => {}
//         },
//         Event::Mouse(_) => {}
//         Event::Resize(_, _) => {}
//     }
//     Ok(())
// }
