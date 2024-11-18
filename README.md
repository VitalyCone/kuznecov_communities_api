<h1>Kuznecov Communities API<h1>
<h5>
<p><b>FOR USING THIS API YOU MUST RUN <b><a href="https://github.com/VitalyCone/kuznecov_cloud_storage_api">https://github.com/VitalyCone/kuznecov_cloud_storage_api</a><p>
</p>Usage</p>
<p>1. Clone this repository</p>
<code>git clone https://github.com/VitalyCone/kuznecov_communities_api</code>
<p></p>
<p>1. Build and run <i>docker-compose.yml</i></p>
<code>docker-compose up --build kzcv-communities</code>
<p></p>
<p>2. Apply migrations</p>
<code>migrate -path migrations -database "postgres://admin:admin@localhost:5001/kuznecov_communities?sslmode=disable" up</code>
<p></p>
<p>3. Join in swagger on your browser for check API documentation</p>
<a href="http://localhost:8001/swagger/index.html">http://localhost:8001/swagger/index.html</a>
</h5>
