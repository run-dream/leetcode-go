SELECT d.Name Department, e1.Name Employee, e1.Salary Salary
FROM Employee e1
JOIN Department d 
ON e.DepartmentId = d.Id
WHERE 3 > (
    SELECT COUNT(DISTINCT(e2.Salary))
    FROM Employee e2
    WHERE e2.DepartmentId = e1.DepartmentId
        AND e2.Salary > e1.Salary
) 