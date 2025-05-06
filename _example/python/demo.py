from dataclasses import dataclass
from typing import Generic, TypeVar, Optional, Any, Type, List, Dict

from client import ApiClient

T = TypeVar('T')
A = TypeVar('A')
B = TypeVar('B')
C = TypeVar('C')
P = TypeVar('P')


@dataclass
class Dog:
    """
    Dog 狗
    """
    color: str = None
    id: int = None
    nickname: str = None
    price: int = None


@dataclass
class PageData(Generic[T]):
    """
    PageData 
    """
    entities: List[T] = None
    page: int = None
    size: int = None
    total: int = None


@dataclass
class ApiResult(Generic[T]):
    """
    ApiResult 
    """
    code: str = None
    data: T = None
    err: str = None
    msg: str = None


@dataclass
class Cat:
    """
    Cat 猫
    """
    foods: List[str] = None
    nickname: str = None


@dataclass
class MapResult(Generic[T]):
    """
    MapResult 
    """
    code: str = None
    data: Dict[str, T] = None


@dataclass
class User:
    """
    User 用户
    """
    age: int = None
    birthday: str = None
    cats: List[Cat] = None
    enable: bool = None
    gods: Dict[str, Dog] = None
    id: int = None
    kv: Dict[str, str] = None
    name: str = None
    profile: JsonNode = None
    tags: List[str] = None
    timestamp: int = None


class AnimalController:
    """
    动物接口
    """
    client: ApiClient = None

    def __init__(self, client: ApiClient):
        self.client = client

    def addCat(self, body: Cat) -> Optional[int]:
        """
        响应long
        """

        return self.client.post("/animal/cat/add", int, body)

    def addDog(self, kind: str, body: Dog) -> Optional[Dog]:
        """
        路径+Body
        """

        return self.client.post("/animal/add/%s/dog".format(kind), Dog, body)

    def cats(self, ) -> Optional[List[Cat]]:
        """
        响应list
        """
        params = {}
        return self.client.get("/animal/cats", List[Cat], params)

    def dog(self, id: int) -> Optional[Dog]:
        """
        响应对象
        """
        params = {}
        return self.client.get("/animal/dog/%s".format(id), Dog, params)

    def mapCats(self, ) -> Optional[Dict[str, Cat]]:
        """
        响应map
        """
        params = {}
        return self.client.get("/animal/map/cats", Dict[str, Cat], params)

    def mapDog(self, ) -> Optional[Dog]:
        """
        响应Map泛型
        """
        params = {}
        return self.client.get("/animal/map/dogs", Dog, params)

    def pageDogs(self, ) -> Optional[Dog]:
        """
        响应Array泛型
        """
        params = {}
        return self.client.get("/animal/page/dogs", Dog, params)

    def searchDog(self, color: str, kind: str, name: str) -> Optional[Dog]:
        """
        路径+Query
        """

        params = {"color": color, "kind": kind, "name": name, }
        return self.client.get("/animal/query/%s/dog".format(kind), Dog, params)


class OtherApi:
    """
    其他接口
    """
    client: ApiClient = None

    def __init__(self, client: ApiClient):
        self.client = client

    def get(self, ) -> Optional[str]:
        """
        响应基础类型
        """
        params = {}
        return self.client.get("/other/long", str, params)

    def getContent(self, ) -> Optional[int]:
        """
        响应基础类型
        """
        params = {}
        return self.client.get("/other/string", int, params)

    def multiplePathVariable(self, id: str, type: str) -> Optional[List[Dog]]:
        """
        多路径参数
        """
        params = {}
        return self.client.get("/other/boolean/%s/%s".format(type, id), List[Dog], params)


class UserApi:
    """
    用户接口
    """
    client: ApiClient = None

    def __init__(self, client: ApiClient):
        self.client = client

    def add(self, body: User) -> Optional[int]:
        """
        添加用户
        """

        return self.client.post("/user/add", int, body)

    def delete(self, id: int) -> Optional[bool]:
        """
        删除用户
        """
        params = {}
        return self.client.delete("/user/delete/%s".format(id), bool, params)

    def edit(self, id: int, body: User) -> Optional[int]:
        """
        修改用户
        """

        return self.client.put("/user/edit/%s".format(id), int, body)

    def edit_1(self, id: int, keyword: str) -> Optional[User]:
        """
        查询用户
        """

        params = {"id": id, "keyword": keyword, }
        return self.client.get("/user/detail/%s".format(id), User, params)

    def list(self, ) -> Optional[List[User]]:
        """
        用户列表
        """
        params = {}
        return self.client.get("/user/list", List[User], params)

    def page(self, page: int, size: int) -> Optional[PageData[User]]:
        """
        用户分页
        """

        params = {"page": page, "size": size, }
        return self.client.get("/user/page", PageData[User], params)
