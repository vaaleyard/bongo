use mongodb::bson::doc;
use mongodb::error::Result;

use mongodb::sync::Client;

// pub async fn get_dbs() -> Result<Vec<String>, Box<dyn Error>> {
pub fn get_database_names(mongo_uri: String) -> Result<Vec<String>> {
    let client = Client::with_uri_str(mongo_uri)?;

    client.list_database_names(None, None)
}

pub fn get_collections_from_db(database_name: String, mongo_uri: String) -> Result<Vec<String>> {
    let client = Client::with_uri_str(mongo_uri)?;

    client
        .database(&database_name)
        .list_collection_names(doc! {"type": "collection"})
}

pub fn get_views_from_db(database_name: String, mongo_uri: String) -> Result<Vec<String>> {
    let client = Client::with_uri_str(mongo_uri)?;

    client
        .database(&database_name)
        .list_collection_names(doc! {"type": "view"})
}

pub fn get_users_from_db(database_name: String, mongo_uri: String) -> Result<Vec<String>> {
    let client = Client::with_uri_str(mongo_uri)?;

    let all_users = client
        .database(&database_name)
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
