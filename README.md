# Layered architecture pattern in Golang
[![Icon](./docs/images/banner.png)](https://github.com/yael-castro)

This is a project template for Golang that uses the layered architecture pattern


❌ IMPORTANT ❌

- The following information was recapitulated from the book [Learning Domain-Driven Design, Vlad Khononov](https://www.oreilly.com/library/view/learning-domain-driven-design/9781098100124/)
- It is very important to read this document completely to understand the architecture pattern

###### Architecture style explained
```
internal
├── business    (business logic layer)
├── dependency  (manage dependencies)
├── handler     (presentation layer)
├── model       (data transfer objects, business objects, errors and enums)
└── repository  (data access layer)

📝 The packages that are not a layer, are an extension of this pattern made by me based on my experience
```
##### 1️⃣ Presentation layer
Implements the program's user interface for interactions with its consumers.
In the pattern's original form, this layer denotes a Graphical User Interface, such as a web interface or a desktop application.

In modern systems, however, the presentation layer has a broader scope: that is, all means for triggering the program's behavior,
both synchronous and asynchronous.
For example:
- Graphical user interface (GUI)
- Command-line interface (CLI)
- API for programmatic integration with other systems
- Subscription to events in a message broker
- Message topics for publishing outgoing events

##### 2️⃣ Business logic layer
This is the layer responsible for implementing and encapsulating the program's business logic

##### 3️⃣ Data access layer 
Provides access to persistence mechanisms.
In the pattern's original form, this referred to the system's database. However, as in the case of the presentation layer,
the layer's responsibility is broader for modern systems.

<ol>
    <li>
        Ever since the NoSQL revolution broke out, it is common for a system to work with multiple databases.
        <br>
        For example, a document store can act as the operational database, a search index for dynamic queries, and in-memory database
        for performance-optimized operations.
        <br><br>
    </li>
    <li>
        Traditional databases are not only medium for storing information.
        <br>
        For example, cloud-based object storage can be used to store the system's files, or a message bus can be used to orchestrate
        communication between the program's different functions.
        <br><br>
    </li>
    <li>
        This layer also includes integration with the various external information providers needed to implement the program's functionality:
        APIs provided by the external systems, or cloud vendor's managed services, such as language translation, stock market data, and
        audio transcription
    </li>
</ol>

###### Communication between layers

The layers are integrated in a top-down communication model: each layer can hold dependency only on the layer directly beneath it.

```
     
    |--------------------|           |----------------------|           |-------------------|
    | Presentation layer | ========> | Business logic layer | ========> | Data access layer |
    |--------------------|           |----------------------|           |-------------------|

```
<hr>
<a href="https://www.flaticon.com/free-icons/stack" title="stack icons">Stack icons created by bukeicon - Flaticon</a>