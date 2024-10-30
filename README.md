# api-students
Projeto de API na linguagem Golang do curso Golang-do-zero.

API como projeto final do curso Golang do Zero da comunidade DEV completo.

Routes:
- GET /students - List all students;
- POST /students - Create student;
- GET /students/:id - Get student info;
- PUT /students/:id - Update student;
- DELETE /students/:id - Delete student
- GET /students?active=<true/false> - List all active/non-active students

Struct student:
- Name string
- CPF int
- Email string
- Age int 
- Active bool