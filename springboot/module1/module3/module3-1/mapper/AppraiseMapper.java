package mapper;

import java.util.List;
import model.Appraise;
import model.AppraiseExample;
import org.apache.ibatis.annotations.Param;
import org.apache.ibatis.session.RowBounds;
import org.springframework.stereotype.Repository;

@Repository
public interface AppraiseMapper {
    /**
     * @mbg.generated generated automatically, do not modify!
     */
    long countByExample(AppraiseExample example);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    int deleteByExample(AppraiseExample example);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    int deleteByPrimaryKey(Integer id);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    int insert(Appraise record);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    int insertSelective(Appraise record);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    List<Appraise> selectByExample(AppraiseExample example, RowBounds rowBounds);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    List<Appraise> selectByExample(AppraiseExample example);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    Appraise selectByPrimaryKeyWithLock(Integer id);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    Appraise selectByPrimaryKey(Integer id);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    int updateByExampleSelective(@Param("record") Appraise record, @Param("example") AppraiseExample example);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    int updateByExample(@Param("record") Appraise record, @Param("example") AppraiseExample example);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    int updateByPrimaryKeySelective(Appraise record);

    /**
     * @mbg.generated generated automatically, do not modify!
     */
    int updateByPrimaryKey(Appraise record);
}