-- name: InfoProjectsScientists :many
SELECT
    s.dni,
    s.nomApels,
    p.id,
    p.nombre
FROM cientificos s
         INNER JOIN asignado_a a ON s.dni = a.cientifico
         INNER JOIN proyectos p ON a.proyecto = p.id;
-- name: NumberProjectsScientist :many


SELECT
    s.dni,
    s.nomApels,
    COUNT(DISTINCT a.proyecto) AS num_proyectos
FROM cientificos s
         INNER JOIN asignado_a a ON s.dni = a.cientifico
GROUP BY s.dni, s.nomApels;

-- name: NumberScientistsProject :many


SELECT
    p.id,
    p.nombre,
    COUNT(DISTINCT a.cientifico) AS num_cientificos
FROM proyectos p
         INNER JOIN asignado_a a ON p.id = a.proyecto
GROUP BY p.id, p.nombre;

-- name: NumberHourScientificProject :many
SELECT
    s.dni,
    s.nomApels,
    SUM(a.dedicacion) AS total_dedicacion
FROM cientificos s
         INNER JOIN asignado_a a ON s.dni = a.cientifico
GROUP BY s.dni, s.nomApels;

-- name: MoreEightyHoursProject :many

SELECT
    s.dni,
    s.nomApels
FROM cientificos s
         INNER JOIN asignado_a a ON s.dni = a.cientifico
GROUP BY s.dni, s.nomApels
HAVING COUNT(DISTINCT a.proyecto) > 1 AND AVG(a.dedicacion) > 80;