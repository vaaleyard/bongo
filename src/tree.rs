pub struct User {
    pub user: String,
}

pub struct Collection {
    pub collection: String,
}

pub struct View {
    pub view: String,
}

#[derive(Default)]
pub struct Database {
    pub collections: Vec<Collection>,
    pub views: Vec<View>,
    pub users: Vec<User>,
}

impl Database {
    pub fn new_collection(&mut self, collection: String) {
        self.collections.push(Collection { collection });
    }

    pub fn new_view(&mut self, view: String) {
        self.views.push(View { view });
    }

    pub fn new_user(&mut self, user: String) {
        self.users.push(User { user });
    }
}
