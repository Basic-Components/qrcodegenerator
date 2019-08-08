bash make.sh linux64
source version
docker build -t hsz1273327/qrcodegenerator:$PROJECT_VERSION -t hsz1273327/qrcodegenerator:latest .
docker push hsz1273327/qrcodegenerator
git add .
git commit -m "update to v$PROJECT_VERSION"
git push
git tag v$PROJECT_VERSION
git push origin --tags