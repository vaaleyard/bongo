use mongodb::{
    options::{ClientOptions, ResolverConfig},
    Client,
};
use std::env;
use std::error::Error;

pub async fn get_dbs() -> Result<Vec<String>, Box<dyn Error>> {
    // Load the MongoDB connection string from an environment variable:
    let client_uri =
        env::var("MONGODB_URI").expect("You must set the MONGODB_URI environment var!");
    // A Client is needed to connect to MongoDB:
    // An extra line of code to work around a DNS issue on Windows:
    let options =
        ClientOptions::parse_with_resolver_config(&client_uri, ResolverConfig::cloudflare())
            .await?;
    let client = Client::with_options(options)?;
    // Print the databases in our MongoDB cluster:
    println!("Databases:");
    for name in client.list_database_names(None, None).await? {
        println!("- {}", name);
    }
    let databases = client.list_database_names(None, None).await?;

    Ok(databases)
}
