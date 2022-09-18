docker build -t postgres-data-visualization-image ./

docker run -d --name postgres-data-visualization-container -p 5432:5432 postgres-data-visualization-image