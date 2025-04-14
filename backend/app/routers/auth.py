from fastapi import APIRouter, Request
from wechatpy.oauth import WeChatOAuth
from app.core.config import settings

router = APIRouter(prefix="/auth", tags=["authentication"])

@router.get("/wechat/login")
async def wechat_login():
    oauth = WeChatOAuth(
        app_id=settings.WECHAT_APP_ID,
        secret=settings.WECHAT_APP_SECRET,
        redirect_uri=settings.WECHAT_REDIRECT_URI,
        scope=settings.WECHAT_SCOPE
    )
    auth_url = oauth.authorize_url(state=settings.WECHAT_STATE)
    return {"auth_url": auth_url}

@router.get("/wechat/callback")
async def wechat_callback(code: str, state: str):
    if state != settings.WECHAT_STATE:
        return {"error": "Invalid state"}
    
    oauth = WeChatOAuth(
        app_id=settings.WECHAT_APP_ID,
        secret=settings.WECHAT_APP_SECRET,
        redirect_uri=settings.WECHAT_REDIRECT_URI
    )
    try:
        access_info = oauth.fetch_access_token(code)
        user_info = oauth.get_user_info()
        return {"user_info": user_info}
    except Exception as e:
        return {"error": str(e)}