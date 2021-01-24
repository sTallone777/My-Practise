package mapper;

import java.util.List;
import model.Adjustsalary;
import model.AdjustsalaryExample;
import org.apache.ibatis.annotations.Param;
import org.apache.ibatis.session.RowBounds;
import org.springframework.stereotype.Repository;

@Repository
public interface AdjustsalaryMapper {
    /**
     * @mbg.generated generated automatically, do not modify!
     */
    long countByExample(AdjustsalaryExample example);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    int deleteByExample(AdjustsalaryExample example);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    int deleteByPrimaryKey(Integer id);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    int insert(Adjustsalary record);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    int insertSelective(Adjustsalary record);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    List<Adjustsalary> selectByExample(AdjustsalaryExample example, RowBounds rowBounds);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    List<Adjustsalary> selectByExample(AdjustsalaryExample example);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    Adjustsalary selectByPrimaryKeyWithLock(Integer id);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    Adjustsalary selectByPrimaryKey(Integer id);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    int updateByExampleSelective(@Param("record") Adjustsalary record, @Param("example") AdjustsalaryExample example);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    int updateByExample(@Param("record") Adjustsalary record, @Param("example") AdjustsalaryExample example);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    int updateByPrimaryKeySelective(Adjustsalary record);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    int updateByPrimaryKey(Adjustsalary record);
}