import { Router } from "express";
import { router as productsRouter } from "./modules/products/products-handler"

const collectorRouter = Router();

collectorRouter.use("/", productsRouter)

export { collectorRouter }