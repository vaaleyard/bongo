use mongodb::bson::doc;
use mongodb::error::Result;

use mongodb::sync::Client;
use std::env;

// pub async fn get_dbs() -> Result<Vec<String>, Box<dyn Error>> {
pub fn get_database_names() -> Result<Vec<String>> {
    // Load the MongoDB connection string from an environment variable:
    let client_uri =
        env::var("MONGODB_URI").expect("You must set the MONGODB_URI environment var!");
    // A Client is needed to connect to MongoDB:
    // An extra line of code to work around a DNS issue on Windows:
    let client = Client::with_uri_str(client_uri)?;

    client.list_database_names(None, None)
}

pub fn get_collections_from_db(database: String) -> Result<Vec<String>> {
    let client_uri =
        env::var("MONGODB_URI").expect("You must set the MONGODB_URI environment var!");

    let client = Client::with_uri_str(client_uri)?;

    client
        .database(&database)
        .list_collection_names(doc! {"type": "collection"})
}

pub fn get_views_from_db(database: String) -> Result<Vec<String>> {
    let client_uri =
        env::var("MONGODB_URI").expect("You must set the MONGODB_URI environment var!");

    let client = Client::with_uri_str(client_uri)?;

    client
        .database(&database)
        .list_collection_names(doc! {"type": "view"})
}

pub fn get_users_from_db(database: String) -> Result<Vec<String>> {
    let client_uri =
        env::var("MONGODB_URI").expect("You must set the MONGODB_URI environment var!");

    let client = Client::with_uri_str(client_uri)?;

    let all_users = client
        .database(&database)
        .run_command(doc! {"usersInfo": 1}, None)
        .unwrap();

    let mut users: Vec<String> = Vec::new();

    all_users
        .get_array("users")
        .unwrap()
        .iter()
        .for_each(|user| {
            users.push(user.as_document().unwrap().get("user").unwrap().to_string());
        });

    Ok(users)
}
