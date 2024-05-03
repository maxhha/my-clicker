import os
import requests
import pytest
from concurrent.futures import ThreadPoolExecutor

@pytest.fixture
def base_url():
    return os.environ['CLICKER_URL']

def test_single_thread_clicking(base_url):
    N = 1000
    resp = requests.post(f"{base_url}/api/link/create", json={"redirect": "http://localhost/some/sample"})
    assert resp.status_code == 200
    resp_data = resp.json()
    short_url, link_id = resp_data['url'], resp_data['link_id']

    for i in range(N):
        resp = requests.get(short_url, allow_redirects=False)
        assert resp.status_code == 302

    resp = requests.get(f"{base_url}/api/link/{link_id}/counter")
    assert resp.status_code == 200
    resp_data = resp.json()
    assert resp.json()['total'] == N


def test_multithread_clicking(base_url):
    N = 1000
    WORKERS = 4
    resp = requests.post(f"{base_url}/api/link/create", json={"redirect": "http://localhost/some/sample"})
    assert resp.status_code == 200
    resp_data = resp.json()
    short_url, link_id = resp_data['url'], resp_data['link_id']

    def worker(*args):
        resp = requests.get(short_url, allow_redirects=False)
        assert resp.status_code == 302

    with ThreadPoolExecutor(WORKERS) as executor:
        executor.map(worker, range(N))

    resp = requests.get(f"{base_url}/api/link/{link_id}/counter")
    assert resp.status_code == 200
    resp_data = resp.json()
    assert resp.json()['total'] == N
