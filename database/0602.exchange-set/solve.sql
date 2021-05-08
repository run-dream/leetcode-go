--- region using if

select
if(id < (select count(*) from seat), if(id % 2=0, id-1, id+1), if(id % 2=0, id-1, id)) as id, student
from seat
order by id asc
--- end region

-- using union
SELECT s1.id - 1 as id, s1.student
FROM seat s1
WHERE s1.id MOD 2 = 0

UNION

SELECT s2.id + 1 as id, s2.student
FROM seat s2
WHERE s2.id MOD 2 = 1 AND s2.id != (SELECT MAX(id) FROM Seat)

UNION


SELECT s3.id, s3.student
FROM seat s3
WHERE s3.id MOD 2 = 1 AND s3.id = (SELECT MAX(id) FROM Seat)


ORDER BY id ASC;
--- end region