/* libzlintpq - run zlint from a PostgreSQL function
 * Written by Rob Stradling
 * Copyright (C) 2017 COMODO CA Limited
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

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
