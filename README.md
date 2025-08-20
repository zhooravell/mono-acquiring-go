Monobank Acquiring GoLang SDK
=============================
[![GO CI](https://github.com/zhooravell/mono-acquiring-go/actions/workflows/go.yml/badge.svg)](https://github.com/zhooravell/mono-acquiring-go/actions/workflows/go.yml)

## Supported API Methods

| Description                             | HTTP Method | Endpoint                                                    | Function               |
|-----------------------------------------|-------------|-------------------------------------------------------------|------------------------|
| Створення рахунку                       | POST        | `/api/merchant/invoice/create`                              | CreateInvoice()        |
| Статус рахунку                          | GET         | `/api/merchant/invoice/status?invoiceId={invoiceId}`        | GetInvoiceStatus()     |
| Скасування оплати                       | POST        | `/api/merchant/invoice/cancel`                              | CancelInvoice()        |
| Інвалідація рахунку                     | POST        | `/api/merchant/invoice/remove`                              | RemoveInvoice()        |
| Відкритий ключ                          | GET         | `/api/merchant/pubkey`                                      | GetPublicKey()         |
| Фіналізація суми холду                  | POST        | `/api/merchant/invoice/finalize`                            | FinalizeHold()         |
| Інформація про QR-касу                  | GET         | `/api/merchant/qr/details?qrId={qrId}`                      | GetQRDetails()         |
| Видалення суми оплати QR                | POST        | `/api/merchant/qr/reset-amount`                             | QrResetAmount()        |
| Список QR-кас                           | GET         | `/api/merchant/qr/list`                                     | GetQRList()            | 
| Дані мерчанта                           | GET         | `/api/merchant/details`                                     | GetMerchantDetails()   |
| Виписка за період                       | GET         | `/api/merchant/statement`                                   | GetStatement()         |
| Видалення токенізованої картки          | DELETE      | `/api/merchant/wallet/card`                                 | RemoveWalletCard()     |
| Список карток у гаманці                 | GET         | `/api/merchant/wallet`                                      | GetWalletCardList()    |
| Оплата по токену                        | POST        | `/api/merchant/wallet/payment`                              |                        |
| Оплата за реквізитами                   | POST        | `/api/merchant/invoice/payment-direct`                      |                        |
| Список субмерчантів                     | GET         | `/api/merchant/submerchant/list`                            | GetSubMerchantList()   |
| Квитанція                               | GET         | `/api/merchant/invoice/receipt?invoiceId={invoiceId}`       | GetReceipt()           |
| Фіскальні чеки                          | GET         | `/api/merchant/invoice/fiscal-checks?invoiceId={invoiceId}` | GetFiscalChecks()      |
| Синхронна оплата                        | POST        | `/api/merchant/invoice/sync-payment`                        | SyncPayment()          |
| Список співробітників                   | GET         | `/api/merchant/employee/list`                               | GetEmployeeList()      |
| Список отримувачів розщеплених платежів | GET         | `/api/merchant/split-receiver/list`                         | GetSplitReceiverList() |

## Source(s)

* [Monobank Acquiring](https://monobank.ua/api-docs)
* [ISO 4217](https://www.iso.org/iso-4217-currency-codes.html)
* [ISO 3166-1](https://www.iso.org/iso-3166-country-codes.html)