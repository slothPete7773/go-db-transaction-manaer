Summary 

- View repository as domain-driven's Aggregate. An aggregate is an entity, composed of more granular SQL tables, that related to the specific domain. 

- Anti-pattern: One repository per database table

Don’t create a repository for each database table. Instead, think of the data that needs to be transactionally stored together.
