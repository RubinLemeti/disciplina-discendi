import { generateMongoDbClient } from "../../helper/mongoDbConnection";
import { Db, Collection } from "mongodb";


export class ProductRepository {

    // Promise<number | [ProductDatabaseModel]>
    async getProductCollectionFromDb(limit?: number, offset?: number): Promise<[any[], number]> {
        limit = (!limit) ? 10 : limit
        offset = (!offset) ? 0 : offset

        try {
            const client = await generateMongoDbClient()
            const db: Db = client.db(process.env.DB_NAME);
            const collectionName = process.env.DB_COLLECTION
            if (collectionName == undefined) {
                throw new Error("Collection variable is not defined")
            }

            const collection: Collection = db.collection(collectionName);

            const cursor = collection.find().limit(limit).skip(offset)
            const productsCollection = []
            for await (const doc of cursor) productsCollection.push(doc);

            const total: number = await collection.countDocuments()
            return [productsCollection, total]
        } catch (error) {
            throw error
        }
    }

    //  Promise<ProductDatabaseModel>
    async getProductResourceFromDb() { }

    async addProductResourceInDb() { }

    async updateProductResourceInDb() { }

    async deleteProductResourceInDb() { }

}
