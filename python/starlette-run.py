from starlette.responses import PlainTextResponse
from starlette.applications import Starlette
from starlette.routing import Route


async def homepage(request):
    return PlainTextResponse("Hello, World (python)!")


app = Starlette(debug=False, routes=[
    Route('/', homepage),
])