# customer-identity

Sign-in and the customer profile for Tidewell Mutual. Every customer-facing app gets its tokens here, and every app that needs a customer's name, contact details or how long they have been with us asks here.

Owned by **Customer Platform**. On-call: `#tw-architecture`.

| Contract | Kind | Consumers |
| --- | --- | --- |
| `POST /v1/auth/token` | REST | customer-portal |
| `GET /v1/customers/{id}` | REST | claims-intake, billing-service, customer-portal |
| `customer.profile.updated` | Event | policy-admin |

`GET /v1/customers/{id}` returns `customer_since` and `tenure_years`: whole years since the customer's first policy started, even if they had a gap of under 90 days. See [ADR-0004](docs/adr/0004-tenure.md).
