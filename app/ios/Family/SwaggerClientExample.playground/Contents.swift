import RxSwift
import SwaggerClient

import PlaygroundSupport
PlaygroundPage.current.needsIndefiniteExecution = true

SwaggerClientAPI.basePath = "https://family-staging.taeho.io"

let email = "taeho@taeho.io"
let password = "1234"

public enum ErrorTest : Error {
    case error(Int, Data?, Error)
}

let req = AccountsLogInRequest(authType: AccountsAuthType.email, email: email, password: password)

AccountsServiceAPI.logIn(body: req).subscribe(onNext: { (resp) in
    print("resp")
    print(resp)
    print(resp.accountId!)
}, onError: { (err) in
    print("err")
    print(err)
    print(type(of: err))
    if let errorResp = err as? ErrorResponse {
        switch errorResp {
        case .error(401, _, _): print("401 Unauthroized")
        default: print(err)
        }
    }
    print(err.localizedDescription)
}, onCompleted: {
    print("completed")
})
