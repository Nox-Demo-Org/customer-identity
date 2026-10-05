# ADR-0004: How we count tenure

Status: Accepted (2024-06)

Marketing, pricing and the contact centre each counted "years with us" differently. customer-identity now owns one definition: whole years since the start of the customer's first policy, where a gap of under 90 days between policies does not reset the count. It is returned as `tenure_years` on `GET /v1/customers/{id}`. Other services must not work it out themselves.
