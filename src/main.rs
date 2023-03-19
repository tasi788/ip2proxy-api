#[macro_use] extern crate rocket;

use ip2location::{error, Record, DB, ProxyRecord, Proxy};
use rocket::serde::{Serialize, json::Json};


const FILES: &str = "src/ip2proxy.BIN";

fn ip_lookup_in_ipv4bin(ip: &str) -> Result<ProxyRecord, error::Error> {
    let mut db = DB::from_file(FILES)?;
    let record = db.ip_lookup(ip.parse().unwrap())?;
    match record {
        Record::ProxyDb(proxy) => {
            Ok(proxy)
        },
        _ => Err(error::Error::RecordNotFound),
    }
}


#[derive(Serialize)]
#[serde(crate = "rocket::serde")]
struct Resp {
    status: bool,
    proxy: bool,
    proxy_type: String,
    country: String,
}


#[get("/query?<ip>")]
fn query(ip: &str) -> Option<Json<Resp>> {
    match ip_lookup_in_ipv4bin(ip) {
        Ok(record) => {
            let is_proxy = !matches!(record.is_proxy, Some(Proxy::IsNotAProxy | Proxy::IsAnError));
            let resp = Resp {
                status: true,
                proxy: is_proxy,
                proxy_type: record.proxy_type.unwrap(),
                country: record.country.unwrap().short_name,
            };
            // println!("{:?}", resp);
            return Some(Json(resp));
        }
        Err(e) => {
            let resp = Resp {
                status: false,
                proxy: false,
                proxy_type: "Record Not Found!".to_string(),
                country: "".to_string(),
            };
            return Some(Json(resp));
        }
    }
}

#[rocket::main]
async fn main() -> Result<(), rocket::Error> {
    let _rocket = rocket::build()
        .mount("/", routes![query])
        .launch()
        .await?;

    Ok(())
}