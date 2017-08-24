CREATE OR REPLACE FUNCTION zlint_embedded(
	cert					bytea
) RETURNS SETOF text
AS $$
DECLARE
BEGIN
	RETURN QUERY SELECT unnest(string_to_array(RTRIM(zlint_wrapper(encode(cert, 'base64')), CHR(10)), CHR(10)));
END;
$$ LANGUAGE plpgsql STRICT;
