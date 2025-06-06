from typing import Generic, TypeVar, Optional, Any, Type
import requests
from dataclasses import asdict

T = TypeVar('T')
P = TypeVar('P')


class ApiClient:

    def get(self, path: str, mode: Type[T], params: P = None) -> Optional[T]:
        ...

    def post(self, path: str, mode: Type[T], params: P = None) -> Optional[T]:
        ...

    def put(self, path: str, mode: Type[T], params: P = None) -> Optional[T]:
        ...

    def delete(self, path: str, mode: Type[T], params: P = None) -> Optional[T]:
        ...


class HttClient(ApiClient):
    BASE_URL = "http://localhost:8083/v3/api-docs"

    def get(self, path: str, mode: Type[T], params: P = None) -> Optional[T]:
        url = f"{self.BASE_URL}{path}"
        return requests.get(url=url, params=params).json()

    def post(self, path: str, mode: Type[T], params: P = None) -> Optional[T]:
        url = f"{self.BASE_URL}{path}"
        return requests.post(url=url, json=asdict(params)).json()

    def put(self, path: str, mode: Type[T], params: P = None) -> Optional[T]:
        url = f"{self.BASE_URL}{path}"
        return requests.put(url=url, json=asdict(params)).json()

    def delete(self, path: str, mode: Type[T], params: P = None) -> Optional[T]:
        url = f"{self.BASE_URL}{path}"
        return requests.delete(url=url, json=asdict(params)).json()
