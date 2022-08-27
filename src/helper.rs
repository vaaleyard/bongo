use std::io;

use crossterm::terminal::LeaveAlternateScreen;

pub fn reset_terminal() -> io::Result<()> {
    crossterm::terminal::disable_raw_mode()?;
    crossterm::execute!(io::stdout(), LeaveAlternateScreen)?;

    Ok(())
}
