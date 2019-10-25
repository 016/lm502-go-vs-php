/**
 * 
 * run with Yii 2.0 framework's path console/controllers
 * @author Lee
 *
 */
class ConsoleController extends Controller{

    public function actionDbTest(){

        $maxRun = 50000;

        echo "PHP CLI Db Test $maxRun times! \n\n";

        $startTime = microtime(true);


        $sum = 0;
        $cnt = 0;

        for ($i = 0; $i < $maxRun; $i++){
            $tmpId = mt_rand(1, 30);
            
            //AR
            //$l = AppLocation::find()->where(['l_id'=>$tmpId])->one();

            //DAO
            $sql = 'select * from app_location where l_id = :l_id';
            $CMD = \Yii::$app->db->createCommand($sql);
            $CMD->bindValue(':l_id', $tmpId);
            
            $l = $CMD->queryOne();
            
            //force fake loop
            for ($i = 0; $i < 10; $i++) {
                if ($i%2 == 0) {
                    $sum += $l['l_id'];
                }else{
                    $sum += $l['l_code'];
                }
            }
            
            $cnt ++;
        }



        $endTime = microtime(true);
        echo ">> start $startTime \n";
        echo ">> end $endTime \n";

        $diff = $endTime - $startTime;


        echo "use time : $diff sum: $sum count:$cnt \n\n";


        
        
        
        
    }
}
