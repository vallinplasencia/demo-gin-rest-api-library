source deploy-config.sh # load env
docker-compose -p library -f docker-compose.yml up -d
echo "Wait db service really ready"
sleep 20s
docker start library_library_demo_1