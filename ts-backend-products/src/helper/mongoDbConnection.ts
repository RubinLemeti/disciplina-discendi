import "dotenv/config";
import { MongoClient } from "mongodb";

function generateMongoConnectionString(): string {
    // postgresql://[user[:password]@][netloc][:port][/dbname][?param1=value1&...]
    let user: string | undefined = process.env.DB_USER;
    let password: string | undefined = process.env.DB_PASSWORD;
    let db: string | undefined = process.env.DB_NAME;
    let host: string | undefined = process.env.DB_HOST;
    let port: string | undefined = process.env.DB_PORT;
    let connectionString: string = `mongodb://${user}:${password}@${host}:${port}`;
    //DB_CONN_STRING="mongodb+srv://<username>:<password>@sandbox.jadwj.mongodb.net"

    // console.log(connectionString)
    return connectionString;
}

export async function generateMongoDbClient() {
    try {
        const client: MongoClient = new MongoClient(generateMongoConnectionString());
        await client.connect();
        
        return client;
    } catch (error) {
        throw error
    }
}