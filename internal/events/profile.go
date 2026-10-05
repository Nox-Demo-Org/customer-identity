package events

// TopicProfileUpdated carries {customer_id, changed: ["email"|"phone"|"address"], updated_at}.
// Subscribers today: policy-admin. billing-service and notifications-hub read profiles on demand.
const TopicProfileUpdated = "customer.profile.updated"
