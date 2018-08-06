import PlaygroundSupport
PlaygroundPage.current.needsIndefiniteExecution = true

import RxSwift
import SwaggerClient

var a = "abc"
print(a)

let helloSequence = Observable.just("Hello Rx")

let subscription = helloSequence.subscribe { event in
    print(event)
}

SwaggerClientAPI.basePath = "https://family-staging.taeho.io"
let email = "taeho@taeho.io"
let password = "1234"
let accountsLogInRequest = AccountsLogInRequest(authType: AccountsAuthType.email, email: email, password: password)

AccountsServiceAPI.logIn(body: accountsLogInRequest).subscribe(onNext: { (resp) in
    print("resp")
    print(resp)
    print(resp.accessToken)
    print(resp.accountId)
    print(resp.expiresIn)
    print(resp.refreshToken)
}, onError: { (err) in
    print(err)
    print(err.localizedDescription)
}, onCompleted: {
    print("completed!")
})
