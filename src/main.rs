use ip2location::{error, Record, DB};

const FILES: &str = "src/ip2proxy.BIN";

fn main() {
    ip_lookup_in_ipv4bin().unwrap();
    // println!("Hello, world!");
}

fn ip_lookup_in_ipv4bin() -> Result<(), error::Error> {
    let mut db = DB::from_file(FILES)?;
    let record = db.ip_lookup("104.16.0.132".parse().unwrap())?;
    match record {
        Record::ProxyDb(proxy) => println!("{:?}", proxy),
        _ => println!("No record found"),
    }
    Ok(())
}
