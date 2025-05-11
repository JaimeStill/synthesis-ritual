# CSQL Language Specification

## Overview
CSQL (.csql) is a hybrid language that combines T-SQL with C# functional programming concepts to enable dynamic query generation and parameterization. The language maintains SQL's declarative nature while introducing programming constructs for better expressiveness and reusability.

## File Structure

### SQL Flavor Declaration
Files must begin with a flavor declaration:

```csql
@sql-flavor: sqlserver | postgresql | mysql | sqlite
```

### Imports and Dependencies
Following the flavor declaration, import statements for C# namespaces:

```csql
@using System.Collections.Generic;
@using System.Linq;
```

### State and Parameters
State declarations define the parameters and objects available to the query:

```csql
@state {
    public DateTime StartDate { get; set; }
    public DateTime EndDate { get; set; }
    public List<int> CustomerIds { get; set; }
}
```

### Functions
Functions can be declared to encapsulate common query patterns:

```csql
@function CustomerFilter(List<int> ids) {
    WHERE CustomerId IN @{ids}
}

@function DateRangeFilter(DateTime start, DateTime end) {
    WHERE OrderDate BETWEEN @{start} AND @{end}
}
```

### Main Query Body
The main query using standard T-SQL syntax with interpolation:

```csql
SELECT 
    c.CustomerId,
    c.CustomerName,
    o.OrderDate,
    o.TotalAmount
FROM 
    Customers c
    INNER JOIN Orders o ON c.CustomerId = o.CustomerId
@{CustomerFilter(CustomerIds)}
@{DateRangeFilter(StartDate, EndDate)}
```

## Language Features

### Interpolation
- Simple value interpolation: `@{property}`
- Expression interpolation: `@{property?.ToString() ?? "Unknown"}`
- Collection interpolation: `IN @{collection}`

### Conditional Sections
Sections of SQL that can be conditionally included:

```csql
@if (CustomerIds?.Any() == true) {
    AND CustomerId IN @{CustomerIds}
}
```

### Query Composition
Queries can be composed using functions and subqueries:

```csql
@function PaginatedResults(int page, int pageSize) {
    OFFSET @{(page - 1) * pageSize} ROWS
    FETCH NEXT @{pageSize} ROWS ONLY
}
```

### Type Safety
The language maintains type safety between C# and SQL:
- C# types are automatically converted to appropriate SQL types
- Parameters are strongly typed
- Collection types are properly handled in IN clauses

## Runtime Behavior

### Query Generation
1. Parse .csql file
2. Evaluate state and parameters
3. Process conditional sections
4. Generate final SQL query
5. Create parameter collection for SqlClient

### Execution Model
```csharp
// Usage in C#
var query = CsqlQuery.Load("query.csql");
query.State.StartDate = DateTime.Today.AddDays(-30);
query.State.CustomerIds = new List<int> { 1, 2, 3 };

// Get generated SQL and parameters
(string sql, SqlParameter[] parameters) = query.Generate();

// Execute using standard SqlClient
using var connection = new SqlConnection(connectionString);
using var command = new SqlCommand(sql, connection);
command.Parameters.AddRange(parameters);
```

## Examples

### Complex Query Example

```csql
@sql-flavor: sqlserver
@using System.Linq;

@state {
    public DateTime? StartDate { get; set; }
    public DateTime? EndDate { get; set; }
    public List<int> CategoryIds { get; set; }
    public string SearchTerm { get; set; }
    public int PageSize { get; set; }
    public int PageNumber { get; set; }
}

@function CategoryFilter(List<int> categories) {
    @if (categories?.Any() == true) {
        AND p.CategoryId IN @{categories}
    }
}

@function SearchFilter(string term) {
    @if (!string.IsNullOrEmpty(term)) {
        AND (
            p.ProductName LIKE '%' + @{term} + '%'
            OR p.Description LIKE '%' + @{term} + '%'
        )
    }
}

WITH ProductStats AS (
    SELECT 
        p.ProductId,
        p.ProductName,
        p.CategoryId,
        p.UnitPrice,
        COUNT(od.OrderId) as OrderCount,
        SUM(od.Quantity * od.UnitPrice) as TotalRevenue
    FROM 
        Products p
        LEFT JOIN [Order Details] od ON p.ProductId = od.ProductId
        LEFT JOIN Orders o ON od.OrderId = o.OrderId
    WHERE 
        1=1
        @if (StartDate.HasValue) {
            AND o.OrderDate >= @{StartDate}
        }
        @if (EndDate.HasValue) {
            AND o.OrderDate <= @{EndDate}
        }
        @{CategoryFilter(CategoryIds)}
        @{SearchFilter(SearchTerm)}
    GROUP BY
        p.ProductId,
        p.ProductName,
        p.CategoryId,
        p.UnitPrice
)
SELECT 
    ps.*,
    c.CategoryName
FROM 
    ProductStats ps
    INNER JOIN Categories c ON ps.CategoryId = c.CategoryId
ORDER BY 
    ps.TotalRevenue DESC
OFFSET @{(PageNumber - 1) * PageSize} ROWS
FETCH NEXT @{PageSize} ROWS ONLY
```

## Limitations and Considerations

1. Query Plan Caching
- Dynamic SQL generation may impact query plan caching
- Consider using sp_executesql for better plan reuse
- Implement parameter sniffing hints where needed

2. Security
- All parameters are properly sanitized
- No direct string concatenation for user inputs
- SQL injection prevention built into the parser

3. Performance
- Generated SQL should be monitored for complexity
- Large IN clauses are handled efficiently
- Consider materialized views for complex CTEs

4. Maintainability
- Keep functions focused and well-named
- Document complex conditional logic
- Use consistent formatting for better readability