# Job Tracker App

## Overview

Tracking job apps sucks, wanted to build something to track easier and also learn to code at the same time

## Workflow

Add company that I'm interested in.

Add roles that they're hiring for.

When ready to apply, track lifecycle of the application. Key dates, key contacts, places to take notes during phone screens.

Document and track next steps and key communications needed

View all applications by status

Stretch:

Maybe suggest actions based on the progress of the application?

## Data entities

### Base

- ID
- Created At
- Updated At
- Deleted At

### Company

- Name: string
- Description: string
- Website: string URL
- Industry (tags?) [strings] enum?
- Months - int months
- HQ location - geo location (city, (state), country)
- Open to remote? boolean

### Role

- Title: string
- Company ID: foreign key company
- URL: string
- Technologies: [strings] enum?
- Pay int
- Remote: boolean
- Location: geo location (nullable if remote is true)
- Level: string enum
- Posted on: date

### Application

- Role ID: foreign key role
- Source ID: foreign key source
- Applied on: date
- Actions: [action IDs, foreign key Actions]
- Key contacts: [contact IDs, foreign key Contact]
- Referrer: Contact

### Source

- Name: string
- Type: string enum?
- Website: string url, nullable
- Recruiter: boolean
- Contacts: [contact IDs, foreign key contact]

### Contact

- Name: string
- Email: string
- Phone: string
- Last contacted: date
- First contacted: date

### Action

- Title: string
- Done: boolean

## Resources

- https://github.com/ardanlabs/service
- https://github.com/gorilla/mux
- https://github.com/dgraph-io/dgo#creating-a-client
- https://github.com/dgraph-io/travel/
- https://github.com/ardanlabs/graphql/
