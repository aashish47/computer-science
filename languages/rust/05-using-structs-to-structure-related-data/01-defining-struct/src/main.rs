struct User{
    active:bool,
    username:String,
    email:String,
}

impl User{
    fn inactive (self) -> User{
        User{
            active:false,
            ..self
        }
    }

    fn new (active:bool, username:String, email:String) ->User{
        Self{
            active, 
            username,
            email,
        }
    }
}

struct Color(u32,u32,u32);

fn main() {
    let mut user1 = User::new(
        true,
        String::from("aashish"),
        String::from("aa@gmail.com"),
    );

    user1.email = String::from("a@gmail.com");
    user1.inactive();

    let Color(_x, _y , _z) = Color(1,2,3);
    
}
