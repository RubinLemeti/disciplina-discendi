// src/errors/HttpError.ts
export class HttpError extends Error {
    public statusCode: number;
    public message: string;
  
    constructor(statusCode: number, message: string) {
      super(message);
      this.statusCode = statusCode;
      this.message = message;
  
      // Ensuring that instances of this class are properly serialized when thrown
      Object.setPrototypeOf(this, HttpError.prototype);
    }
  }
  
  // src/errors/ValidationError.ts
  export class ValidationError extends HttpError {
    constructor(message = 'Validation Error') {
      super(422, message);
    }
  }
  
  // src/errors/AuthenticationError.ts
  export class AuthenticationError extends HttpError {
    constructor(message = 'Unauthorized') {
      super(401, message);
    }
  }
  
  // src/errors/AuthorizationError.ts
  export class AuthorizationError extends HttpError {
    constructor(message = 'Forbidden') {
      super(403, message);
    }
  }
  
  // src/errors/InternalServerError.ts
  export class InternalServerError extends HttpError {
    constructor(message = 'Internal Server Error') {
      super(500, message);
    }
  }
  