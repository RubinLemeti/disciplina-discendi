import express, { Application } from "express";
import cors from "cors";
import path from "path";
import { errorHandler } from "./errors/errorHandler";
import { collectorRouter } from "./router"

const app: Application = express();
const port: number = 3000;

// merging the routes
app.use("/", collectorRouter)

// adding the cors middleware
app.use(cors());

// adding the error handler
app.use(errorHandler);

app.get("/", (request, response) => {
    response.statusCode = 200;
    response.send({ response: "This is the API homepage.", code: "200" });
})

app.listen(port, () => {
    console.log(`Products app listening on port ${port}`);
})