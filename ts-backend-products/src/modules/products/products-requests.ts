import Joi from "joi";

export const productsQueryParams = Joi.object({
    limit: Joi.number().greater(0).default(10),
    offset: Joi.number().default(0),
});