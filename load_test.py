import asyncio
import aiohttp
import time

base_url = "https://gobot-backend-ew7jmfmn3a-uw.a.run.app/chat"

async def send_one_request(session, uid):
    async with session.post(base_url+f"/{uid}?level=easy&message=Hello") as resp:
        return resp.status

async def test():
    start = time.time()
    async with aiohttp.ClientSession() as session:
        tasks = [send_one_request(session, uid) for uid in range(1000)]
        results = await asyncio.gather(*tasks)

    print(f"Finished in {time.time() - start} seconds")
    print(results.count(200))

if __name__ == "__main__":
    asyncio.run(test())