import { ProductDatabaseModel } from "./products-models"
import { ProductRepository } from "./products-repository"

export class ProductsService {
    productRepo: ProductRepository

    constructor() {
        this.productRepo = new ProductRepository()
    }

    // Promise<number | [ProductDatabaseModel]>
    public async getProductCollection(limit?: number, offset?: number) {
        return await this.productRepo.getProductCollectionFromDb(limit, offset)
    }

    //  Promise<ProductDatabaseModel>
    public async getProductResource() {
        return await this.productRepo.getProductResourceFromDb()
    }

    public async addProductResource() {
        return await this.productRepo.addProductResourceInDb()
    }

    public async updateProductResource() {
        return await this.productRepo.updateProductResourceInDb()
    }

    public async deleteProductResource() {
        return await this.productRepo.deleteProductResourceInDb()
    }
}
