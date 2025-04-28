import {Dog, OtherApi, User, UserApi} from "./demo";

console.log('Happy developing ✨')


OtherApi.multiplePathVariable("12", "abc").then((e) => {
    console.info(e)
})

let dog: Dog = {color: "", id: 0, nickname: "", price: 0}
let user: User = {
    age: 12,
    birthday: "2023-10-10",
    cats: [],
    enable: false,
    gods: {
        "d1": dog,
        "d2": dog
    },
    id: 0,
    kv: {
        "k1": "v1",
        "k2": "v2"
    },
    name: "dev",
    profile: undefined,
    tags: ["t1", "t2", "t3"],
    timestamp: 120344
}


UserApi.add(user).then(res => {
    console.log(res.data)
})
