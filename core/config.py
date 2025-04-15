from pydantic_settings import BaseSettings

class Settings(BaseSettings):
    WECHAT_APP_ID: str
    WECHAT_APP_SECRET: str
    WECHAT_REDIRECT_URI: str = "http://localhost:51025/auth/wechat/callback"
    WECHAT_SCOPE: str = "snsapi_login"
    WECHAT_STATE: str = "STATE"

    class Config:
        env_file = ".env"

settings = Settings()