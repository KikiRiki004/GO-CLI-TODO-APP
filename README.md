
### GO CLI TODO APP  

**Use Example:**  
```bash
Usage of ./todo:
  -add string  
        Add a new todo specify title  
  -del int  
        Specify todo by index to delete (default -1)  
  -edit string  
        Edit a todo by index & specify a new title. id:new_title  
  -list  
        List all todos  
  -toggle int  
        Specify todo by index to toggle complete true/false (default -1)
  
todo_cli % ./todo -list  
┌───┬──────────────────────┬───────────┬───────────────────────────────┬──────────────┐  
│ # │        Title         │ Completed │          Created At           │ Completed At │  
├───┼──────────────────────┼───────────┼───────────────────────────────┼──────────────┤  
│ 0 │ Go Gym               │ ❌        │ Wed, 18 Dec 2024 16:09:25 CET │              │  
│ 1 │ Finish todo code     │ ❌        │ Wed, 18 Dec 2024 16:09:55 CET │              │  
│ 2 │ Read book            │ ❌        │ Wed, 18 Dec 2024 16:10:16 CET │              │  
│ 3 │ Understand todo code │ ❌        │ Wed, 18 Dec 2024 16:10:31 CET │              │  
└───┴──────────────────────┴───────────┴───────────────────────────────┴──────────────┘  
todo_cli % ./todo -toggle 1  
todo_cli % ./todo -toggle 3  
todo_cli % ./todo -list  
┌───┬──────────────────────┬───────────┬───────────────────────────────┬───────────────────────────────┐  
│ # │        Title         │ Completed │          Created At           │         Completed At          │  
├───┼──────────────────────┼───────────┼───────────────────────────────┼───────────────────────────────┤  
│ 0 │ Go Gym               │ ❌        │ Wed, 18 Dec 2024 16:09:25 CET │                               │  
│ 1 │ Finish todo code     │ ✅        │ Wed, 18 Dec 2024 16:09:55 CET │ Wed, 18 Dec 2024 16:11:07 CET │  
│ 2 │ Read book            │ ❌        │ Wed, 18 Dec 2024 16:10:16 CET │                               │  
│ 3 │ Understand todo code │ ✅        │ Wed, 18 Dec 2024 16:10:31 CET │ Wed, 18 Dec 2024 16:11:13 CET │  
└───┴──────────────────────┴───────────┴───────────────────────────────┴───────────────────────────────┘  
todo_cli % go run ./ -list  
┌───┬──────────────────────┬───────────┬───────────────────────────────┬───────────────────────────────┐  
│ # │        Title         │ Completed │          Created At           │         Completed At          │  
├───┼──────────────────────┼───────────┼───────────────────────────────┼───────────────────────────────┤  
│ 0 │ Go Gym               │ ❌        │ Wed, 18 Dec 2024 16:09:25 CET │                               │  
│ 1 │ Finish todo code     │ ✅        │ Wed, 18 Dec 2024 16:09:55 CET │ Wed, 18 Dec 2024 16:11:07 CET │  
│ 2 │ Read book            │ ❌        │ Wed, 18 Dec 2024 16:10:16 CET │                               │  
│ 3 │ Understand todo code │ ✅        │ Wed, 18 Dec 2024 16:10:31 CET │ Wed, 18 Dec 2024 16:11:13 CET │  
└───┴──────────────────────┴───────────┴───────────────────────────────┴───────────────────────────────┘  
