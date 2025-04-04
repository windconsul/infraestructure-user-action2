inputs:
domain:
description: 'Target domain (e.g., un.nuevo.dominio.com)'
required: true
instance:
description: 'Instance ID (e.g., 67200dac0502f3ba600968ce)'
required: true
alternativeName:
description: 'Alternative name (e.g., alternativo-com)'
required: false

outputs:
variables de entorno
INFRAESTRUCTURE_USER_ACTION_USER_ID
INFRAESTRUCTURE_USER_ACTION_INSTANCE_ID
INFRAESTRUCTURE_USER_ACTION_PASSWORD
INFRAESTRUCTURE_USER_ACTION_USERNAME
INFRAESTRUCTURE_USER_ACTION_PRIVATE_KEY

Tags     Instancias
         ecs-wind-001 - ecs-registry
test     ecs-wind-002 - ecs-SportTest
         ecs-wind-003 - ecs-MySQL-0001
         ecs-wind-004 - ecs-BoundaryController
prod     ecs-wind-005 - ecs-PageWind
         ecs-wind-006 - ecs-RTWind
prod     ecs-wind-007 - ecs-CRMWind
stagging ecs-wind-008 - ecs-SistgoWind
