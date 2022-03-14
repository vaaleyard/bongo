use crate::{app::App, mongo};

use futures::TryFutureExt;
use tui::{
    backend::Backend,
    layout::{Alignment, Constraint, Direction, Layout},
    style::{Color, Modifier, Style},
    text::{Span, Spans},
    widgets::{Block, Borders, Paragraph},
    Frame,
};

pub fn ui<B: Backend>(f: &mut Frame<B>, _app: &App) {
    let size = f.size();

    // Words made "loooong" to demonstrate line breaking.
    let s = "Very long string.\n";
    let mut long_line = s.repeat(usize::from(size.width) / s.len() + 4);
    long_line.push('\n');

    let block = Block::default().style(Style::default().bg(Color::Black).fg(Color::White));
    f.render_widget(block, size);

    let chunks = Layout::default()
        .direction(Direction::Horizontal)
        .margin(2)
        .constraints([Constraint::Percentage(20), Constraint::Percentage(80)].as_ref())
        .split(size);

    let create_block = |title| {
        Block::default()
            .borders(Borders::ALL)
            .style(Style::default().bg(Color::Black).fg(Color::White))
            .title(Span::styled(
                title,
                Style::default().add_modifier(Modifier::BOLD),
            ))
    };

    let db_tmp = mongo::get_dbs();
    // println!("{:?}", db_tmp);

    let dbs = vec!["admin", "config", "local", "bergo"];
    let bergo: Vec<Spans> = dbs
        .iter()
        .map(|db| {
            Spans::from(vec![Span::styled(
                format!(" ▶ {}", db.clone()),
                Style::default().bg(Color::Black).fg(Color::White),
            )])
        })
        .collect();

    let paragraph = Paragraph::new(bergo.clone())
        .style(Style::default().bg(Color::White).fg(Color::Black))
        .block(create_block("Databases"))
        .alignment(Alignment::Left);
    f.render_widget(paragraph, chunks[0]);
}
