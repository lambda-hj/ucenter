from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from app.routers import auth

app = FastAPI()

# CORS configuration for mobile app
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Include routers
app.include_router(auth.router)

@app.get('/')
async def root():
    return {'message': 'Hello World'}
