Monobank Acquiring GoLang SDK
=============================

## Supported API Methods

| Description                             | HTTP Method | Endpoint                                                    | Function           |
|-----------------------------------------|-------------|-------------------------------------------------------------|--------------------|
| Створення рахунку                       | POST        | `/api/merchant/invoice/create`                              | CreateInvoice()    |
| Статус рахунку                          | GET         | `/api/merchant/invoice/status?invoiceId={invoiceId}`        | GetInvoiceStatus() |
| Скасування оплати                       | POST        | `/api/merchant/invoice/cancel`                              | CancelInvoice()    |
| Інвалідація рахунку                     | POST        | `/api/merchant/invoice/remove`                              |                    |
| Відкритий ключ                          | GET         | `/api/merchant/pubkey`                                      |                    |
| Фіналізація суми холду                  | POST        | `/api/merchant/invoice/finalize`                            |                    |
| Інформація про QR-касу                  | GET         | `/api/merchant/qr/details?qrId={qrId}`                      |                    |
| Видалення суми оплати QR                | POST        | `/api/merchant/qr/reset-amount`                             |                    |
| Список QR-кас                           | GET         | `/api/merchant/qr/list`                                     |                    | 
| Дані мерчанта                           | GET         | `/api/merchant/details`                                     |                    |
| Виписка за період                       | GET         | `/api/merchant/statement`                                   |                    |
| Видалення токенізованої картки          | DELETE      | `/api/merchant/wallet/card`                                 |                    |
| Список карток у гаманці                 | GET         | `/api/merchant/wallet`                                      |                    |
| Оплата по токену                        | POST        | `/api/merchant/wallet/payment`                              |                    |
| Оплата за реквізитами                   | POST        | `/api/merchant/invoice/payment-direct`                      |                    |
| Список субмерчантів                     | GET         | `/api/merchant/submerchant/list`                            |                    |
| Квитанція                               | GET         | `/api/merchant/invoice/receipt?invoiceId={invoiceId}`       |                    |
| Фіскальні чеки                          | GET         | `/api/merchant/invoice/fiscal-checks?invoiceId={invoiceId}` |                    |
| Синхронна оплата                        | POST        | `/api/merchant/invoice/sync-payment`                        |                    |
| Список співробітників                   | GET         | `/api/merchant/employee/list`                               |                    |
| Список отримувачів розщеплених платежів | GET         | `/api/merchant/split-receiver/list`                         |                    |

## Source(s)

* [Monobank Acquiring](https://monobank.ua/api-docs)
* [ISO 4217](https://www.iso.org/iso-4217-currency-codes.html)
* [ISO 3166-1](https://www.iso.org/iso-3166-country-codes.html)