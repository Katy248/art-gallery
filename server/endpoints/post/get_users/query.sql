
	SELECT 
		p.id
		, p.description
		, p.image_url
		, p.publisher_id
		, u.name
		, CAST(CASE WHEN ps.user_id IS NULL THEN 0 ELSE 1 END AS BOOLEAN) as saved
	FROM 
		posts p 
			LEFT JOIN users u ON p.publisher_id = u.id
			LEFT JOIN post_saves ps ON 
				p.id = ps.post_id 
				and ps.user_id = 4

	WHERE publisher_id = 3 and (SELECT COUNT(*) FROM post_saves WHERE post_id = p.id and user_id = 3) <> 0
	
	ORDER BY 
		p.created_at DESC
