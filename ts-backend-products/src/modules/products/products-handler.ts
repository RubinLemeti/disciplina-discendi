import { Router, Request, Response, json } from "express";
import { CollectionJsonStandardResponse, ErrorJsonStandardResponse, SuccessJsonStandardResponse } from "../../helper/genericJsonResponses";
import { logger } from "../../helper/logger";
import { ProductsService } from "./products-service";
import { Paginator } from "../../paginator";
import { productsQueryParams } from "./products-requests";

const router = Router()

router.get("/products", getProductCollection)
async function getProductCollection(request: Request, response: Response) {
    try {
        const productSer = new ProductsService()
        const [productsCollection, total] = await productSer.getProductCollection(undefined, undefined)

        const validatedQueryParams = await productsQueryParams.validateAsync({
            limit: request.query?.limit,
            offset: request.query?.offset,
        });

        const paginatorObj = new Paginator(
            validatedQueryParams.limit,
            validatedQueryParams.offset,
            productsCollection.length,
            total,
            "/products"
        )
        const paginatedCollection = paginatorObj.paginate(productsCollection)

        logger.info("Success")

        CollectionJsonStandardResponse(response, '/products', paginatedCollection.data, paginatedCollection.meta, undefined)

        // response.status(200).send(productsCollection)
    } catch (error) {
        logger.error(error)
        ErrorJsonStandardResponse(response, "/products", undefined, undefined, undefined)
        return;
    }
}

// router.get("/products/:product_id")
// async function getProductResource() { }

router.post("/products", addProductResource)
async function addProductResource() { }

// router.patch("/products")
// async function updateProductResource() { }

// router.delete("/products")
// async function deleteProductResource() { }

export { router }