CREATE OR REPLACE FUNCTION zlint_embedded(
	cert					bytea
) RETURNS SETOF text
AS $$
DECLARE
BEGIN
	RETURN QUERY SELECT *
					FROM (SELECT unnest(string_to_array(RTRIM(zlint_wrapper(encode(cert, 'base64')), CHR(10)), CHR(10))) LINT_ISSUE) sub
					ORDER BY case(substr(sub.LINT_ISSUE, 1, 1))
								WHEN 'F' THEN 1
								WHEN 'E' THEN 2
								WHEN 'W' THEN 3
								ELSE 4
							END,
						substr(sub.LINT_ISSUE, 3);
END;
$$ LANGUAGE plpgsql STRICT;
