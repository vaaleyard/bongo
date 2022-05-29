use mongodb::error::Result;
use mongodb::sync::Client;
use std::env;

// pub async fn get_dbs() -> Result<Vec<String>, Box<dyn Error>> {
pub fn get_dbs() -> Result<Vec<String>> {
    // Load the MongoDB connection string from an environment variable:
    let client_uri =
        env::var("MONGODB_URI").expect("You must set the MONGODB_URI environment var!");
    // A Client is needed to connect to MongoDB:
    // An extra line of code to work around a DNS issue on Windows:
    let client = Client::with_uri_str(client_uri)?;

    client.list_database_names(None, None)
}
