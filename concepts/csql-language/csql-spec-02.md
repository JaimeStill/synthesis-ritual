# CSQL Language Specification

## Overview
CSQL is a hybrid language that combines T-SQL with C# functional programming concepts to enable dynamic query generation and parameterization. The language maintains SQL's declarative nature while introducing programming constructs for better expressiveness and reusability.

## Connection Configuration

### SQL Flavor and Connection Declaration
Files must begin with a connection configuration:

```csql
@sql-config {
    flavor: sqlserver | postgresql | mysql | sqlite,
    // Direct connection string
    connection: "Server=myServerAddress;Database=myDataBase;User Id=myUsername;Password=myPassword;",
    // OR configuration key
    connectionKey: "MyAppDatabase"
}
```

### Imports and Dependencies
Following the configuration, import statements for C# namespaces:

```csql
@using System.Collections.Generic;
@using System.Linq;
```

## Query Definition

### Query Class Structure
Each .csql file generates a C# class that inherits from CsqlQuery:

```csql
@query ProductAnalytics {
    // State properties
    public DateTime StartDate { get; set; }
    public DateTime EndDate { get; set; }
    public List<int> CustomerIds { get; set; }
    
    // Optional constructor
    public ProductAnalytics() {
        StartDate = DateTime.Today.AddDays(-30);
        CustomerIds = new List<int>();
    }
}
```

### Functions
Functions are transformed into protected methods of the generated class:

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

## C# Integration

### Generated Class
The CSQL compiler generates a C# class that can be used directly in code:

```csharp
// Generated class structure
public class ProductAnalytics : CsqlQuery
{
    public DateTime StartDate { get; set; }
    public DateTime EndDate { get; set; }
    public List<int> CustomerIds { get; set; }

    public ProductAnalytics() 
    {
        StartDate = DateTime.Today.AddDays(-30);
        CustomerIds = new List<int>();
    }

    protected string CustomerFilter(List<int> ids) => 
        QueryBuilder.BuildFragment("WHERE CustomerId IN @ids", new { ids });

    protected string DateRangeFilter(DateTime start, DateTime end) =>
        QueryBuilder.BuildFragment("WHERE OrderDate BETWEEN @start AND @end", 
                                 new { start, end });
}
```

### Usage in C#
The query can be used as a regular C# class:

```csharp
// Registration in dependency injection
services.AddCsqlQueries(builder => {
    builder.AddQuery<ProductAnalytics>();
});

// Usage in a service
public class ProductService
{
    private readonly ProductAnalytics _query;

    public ProductService(ProductAnalytics query)
    {
        _query = query;
    }

    public async Task<IEnumerable<Product>> GetProductsAsync()
    {
        _query.StartDate = DateTime.Today.AddDays(-7);
        _query.CustomerIds = new List<int> { 1, 2, 3 };
        
        using var connection = await _query.CreateConnectionAsync();
        return await connection.QueryAsync<Product>(_query);
    }
}
```

## Advanced Features

### Query Result Type Definition
Define the expected result type directly in the query:

```csql
@query ProductSummary<ProductStats> {
    public record ProductStats(
        int ProductId,
        string Name,
        decimal Revenue,
        int OrderCount
    );
    
    // Query state
    public DateTime StartDate { get; set; }
}
```

### Automatic Type Mapping
The compiler generates appropriate Dapper mappings:

```csharp
public class ProductSummary : CsqlQuery<ProductStats>
{
    public record ProductStats(
        int ProductId,
        string Name,
        decimal Revenue,
        int OrderCount
    );
    
    protected override void ConfigureMapping(SqlMapper.ITypeMap mapping)
    {
        mapping.Map<ProductStats>(
            productId: "ProductId",
            name: "Name",
            revenue: "Revenue",
            orderCount: "OrderCount"
        );
    }
}
```

### Connection Management
The base CsqlQuery class handles connection management:

```csharp
public abstract class CsqlQuery
{
    protected IDbConnection Connection { get; private set; }
    
    public virtual async Task<IDbConnection> CreateConnectionAsync()
    {
        // Creates connection based on config
        return await ConnectionFactory.CreateAsync(Configuration);
    }
    
    public virtual async Task<IEnumerable<T>> ExecuteAsync<T>()
    {
        using var connection = await CreateConnectionAsync();
        return await connection.QueryAsync<T>(GenerateQuery(), GetParameters());
    }
}
```

## Complete Example

```csql
@sql-config {
    flavor: sqlserver,
    connectionKey: "NorthwindDb"
}

@query ProductAnalytics<ProductSummary> {
    public record ProductSummary(
        int ProductId,
        string ProductName,
        decimal UnitPrice,
        int UnitsInStock,
        decimal TotalRevenue
    );

    // State
    public DateTime? StartDate { get; set; }
    public List<int> CategoryIds { get; set; }
    public string SearchTerm { get; set; }
    
    public ProductAnalytics()
    {
        StartDate = DateTime.Today.AddDays(-30);
        CategoryIds = new List<int>();
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
    
    SELECT 
        p.ProductId,
        p.ProductName,
        p.UnitPrice,
        p.UnitsInStock,
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
        @{CategoryFilter(CategoryIds)}
        @{SearchFilter(SearchTerm)}
    GROUP BY
        p.ProductId,
        p.ProductName,
        p.UnitPrice,
        p.UnitsInStock
}
```

## Environmental Integration

### Configuration
The query system integrates with .NET's configuration system:

```csharp
// In Program.cs or Startup.cs
services.AddCsqlQueries(options => {
    options.ConnectionStrings = configuration.GetSection("ConnectionStrings");
    options.DefaultProvider = SqlProvider.SqlServer;
});
```

### Dependency Injection
Queries can be registered with different lifetimes:

```csharp
services.AddCsqlQueries(builder => {
    // Transient queries (new instance per request)
    builder.AddTransient<ProductAnalytics>();
    
    // Scoped queries (one instance per scope)
    builder.AddScoped<OrderSummary>();
    
    // Singleton queries (shared instance)
    builder.AddSingleton<CategoryLookup>();
});
```

## Best Practices

1. Query Organization
- One query class per file
- Clear separation of concerns
- Reusable functions for common patterns

2. State Management
- Immutable where possible
- Clear initialization in constructors
- Validation attributes for constraints

3. Connection Handling
- Use configuration keys over direct strings
- Implement retry policies
- Proper connection disposal

4. Performance
- Lazy query generation
- Parameter caching
- Efficient type mapping

5. Security
- Connection string encryption
- Parameter validation
- SQL injection prevention