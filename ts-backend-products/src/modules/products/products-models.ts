export interface ProductDatabaseModel {
    name: string
    description: string
    price: number
    tags: [string]
    inventoryCount: number
    createdAt: string
    updatedAt: string
    createdBy: string
    updatedby: string
}