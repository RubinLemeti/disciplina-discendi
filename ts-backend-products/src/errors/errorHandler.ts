// src/middleware/errorHandler.ts
import { Request, Response, NextFunction } from 'express';
import { HttpError } from './errorClasses';

export const errorHandler = (
  err: HttpError,
  req: Request,
  res: Response,
  next: NextFunction
) => {
  // If error doesn't have a statusCode, default it to 500
  const statusCode = err.statusCode || 500;
  const message = err.message || 'An unexpected error occurred';

  console.error(`[${new Date().toISOString()}] ${statusCode}: ${message}`);

  res.status(statusCode).json({
    error: {
      statusCode,
      message,
    },
  });
  return;
};
